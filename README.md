# Password Police
Password Police is a solution to enforce a single password policy between the front-end and the back-end of an application.

Password Police facilitates sharing of a single passsword policy definition between different parts of an application, even if those parts of the application are written in different programming langauges.

## Why use Password Police?
Let me explain the problem that this project trying to solve. Imagine a simple web project with a front-end written in React, and a back-end written in Go. When a new user signs up on our website, we want them to enter a password. We have some rules that we want the password to comply with, for example a minimum number of characters for the password. Those rules are called a "password policy". We want to check the password policy on the front-end, while the user is typing the password, so we can give the user instant feedback. We also want to check the password policy on the back-end. Why? Because we know that a savvy user can get around our front-end javascript checks. So we also implement the password policy in our back-end code.

Two months down the road our product manager sends us an email stating that our password policy is not strict enough; the minimum password length needs to be longer. So we go into the front-end and change our code. We go into the back-end and change our code. Three months later the security officer sends us an email letting us know that we should enforce the use of a special character in our passwords. Now we have to again change our front-end code and our back-end code. That's a pain.

Now imagine what would happen if you inadvertenly coded a different password policy in the front-end and the back-end. That may turn into a debugging nightmare, especially if the password policy is very verbose.

With Password Police you never have to edit any front-end or back-end code to change the password policy. The password policy is defined in a single file that resides on the back-end. Both the front-end and the back-end use the password policy defined in this one file. Changing the password policy in this file will instantly update the front-end and back-end behavior.

## Status
The golang module is under development. Development of modules/packages/libraries for other languages has not yet started.

## Examples
The [examples](examples/README.md) folder contains example projects that showcase the use of Password Police
