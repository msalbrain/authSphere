Dev features

Database (YES)
Logging (NO)
Notification (NO)
Services (NO)
Testing (NO)
Caching (NO)
Configuration Management (for diff env) (NO)

Documentation (YES)



General Features

general token - for only admin use

email and password -

uniq basic token for each signup

uniq access token (bearer) with refresh token support (optional)

(can be enabled or not) 2fa - after signup verification token or link is sent to mail to verify

/create user

/signup

/login

/refresh token

/endpoint created to accept verify token

/get session endpoint

/user info

/current_user

/update_current_user

/delete user

/password change req

/forgot password

/verify password reset

/link account

/signout



Refferal






Social Login

google
twitter
apple (not sure)
discord
microsoft
github


Admin ['cli', 'web']

View paginated list of users

add and remove resource endpoint

set rules on each resource endpoint

manage users

security settings

Customization:
    Customize the look and feel of the authentication service dashboard to match the branding and design of the application.

User Notifications:
   update email template

analytics:
   email use
   traffic use

logging for administrative doings








