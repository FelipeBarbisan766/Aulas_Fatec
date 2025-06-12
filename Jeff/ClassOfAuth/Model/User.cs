using Microsoft.AspNetCore.Identity;

namespace ClassOfAuth.Model
{
    public class User : IdentityUser
    {
        public string FisrtName { get; set; }
        public string LastName { get; set; }
        public DateTime DateCreat { get; set; } = DateTime.UtcNow;
    }
}
