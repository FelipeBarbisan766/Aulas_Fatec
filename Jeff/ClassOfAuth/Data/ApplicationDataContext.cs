using ClassOfAuth.Model;
using Microsoft.AspNetCore.Identity.EntityFrameworkCore;

namespace ClassOfAuth.Data
{
    public class ApplicationDataContext : IdentityDbContext<User>
    {
    }
}
