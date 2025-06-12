using ClassOfAuth.Model;
using Microsoft.AspNetCore.Identity;
using Microsoft.IdentityModel.Tokens;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;

namespace ClassOfAuth.Service
{
    public class JWTService
    {
        private readonly IConfiguration _config;
        private readonly UserManager<User> _userManeger;
        private readonly SymmetricSecurityKey _jwtKey;

        public JWTService(IConfiguration config, UserManager<User> userManeger, SymmetricSecurityKey _jwtKey)
        {
            _config = config;
            _userManeger = userManeger;
            _jwtKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes("JWT:Key"));
        }
        public async Task<string> CreateJWT(User user)
        {
            var userClaims = new List<Claim>
            {
                new Claim(ClaimTypes.NameIdentifier,user.Id),
                new Claim(ClaimTypes.Email,user.UserName),
                new Claim(ClaimTypes.GivenName, user.FisrtName),
                new Claim(ClaimTypes.Surname, user.LastName)
            };
            var roles = await _userManeger.GetRolesAsync(user);
            userClaims.AddRange(roles.Select(role => new Claim(ClaimTypes.Role, role)));
            
            var credentials = new SigningCredentials(_jwtKey, SecurityAlgorithms.HmacSha256);
            var tokenDescriptor = new SecurityTokenDescriptor
            {
                Subject = new ClaimsIdentity(userClaims),
                Expires = DateTime.UtcNow.AddMinutes(int.Parse(_config["JWT: ExpiresMinutes"])),
                SigningCredentials = credentials,
                Issuer = _config["JWT: Issuer"]
            };
            var tokenHandler = new JwtSecurityTokenHandler();
            var jwt = tokenHandler.CreateToken(tokenDescriptor);
            return tokenHandler.WriteToken(jwt);
        }


    }
}
