using CleanCodeAMS;

var user = new User();
user.Name = "John Doe";
user.Email = "JohnDoe@email.com";

var userAdd = new UserAdd();
userAdd.addUser(user.Name, user.Email);

var email = new Email();
var body = $"Hello {user.Name},\n\nWelcome to our service! We are glad to have you on board.\n\nBest regards,\nThe Team";
email.SendEmail(user.Email,body);

