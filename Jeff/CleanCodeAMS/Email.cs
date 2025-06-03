using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CleanCodeAMS
{
    internal class Email
    {
        public void SendEmail(string email,string body)
        {
            if (string.IsNullOrWhiteSpace(email) || string.IsNullOrWhiteSpace(body))
            {
                throw new ArgumentException("Email and body cannot be empty.");
            }
            // Code to send email
            Console.WriteLine("Email sent successfully.");
            Console.WriteLine($"Sent to {email} this email: {body}.");
        }
    }
}
