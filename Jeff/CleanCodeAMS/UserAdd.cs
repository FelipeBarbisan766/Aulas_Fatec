using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CleanCodeAMS
{
    internal class UserAdd
    {
        public void addUser(string name,string email)//entity User
        {
            if(string.IsNullOrWhiteSpace(name) || string.IsNullOrWhiteSpace(email))
            {
                throw new ArgumentException("Name and email cannot be empty.");
            }
            // Code to add user
            Console.WriteLine("User added successfully.");
            Console.WriteLine($"User Name {name}.");
            Console.WriteLine($"User Email {email}.");
            ValidateTaxId(taxId);
        }

        public void ValidateTaxId(string taxId) {
            if (string.IsNullOrWhiteSpace(taxId))
                throw new ArgumentException("Tax ID cannot be empty.");
            // Example validation: check if the tax ID is a valid format (e.g., 10 digits)
            if (!System.Text.RegularExpressions.Regex.IsMatch(taxId, @"^\d{10}$"))
            {
                throw new ArgumentException("Invalid Tax ID format. It should be 10 digits.");
            }
            Console.WriteLine("Tax ID is valid.");
        }
}
