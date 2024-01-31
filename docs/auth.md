# How the auth system works

## Register

To register send a `POST` request to `/register` with the following body

```json
{
  "email": "gonagutor@gmail.com",
  "password": "Testtest1@",
  "name": "Gonzalo",
  "surname": "Aguado"
}
```

If everything went OK you should receive a 200 code and an email should be sent
_If you get a Gateway error while on local, check if the email credentials are correct_

## Validate email

To validate the email use the generated code sent by email and send a `GET` request to `/validate-email` using the code in a `token` query param
If everything went OK you should receive a 200 code

## Login

To login send a `POST` request to `/login` with the following body

```json
{
  "email": "gonagutor@gmail.com",
  "password": "Testtest1@"
}
```

The response is a little more intricate, apart from the usual `code` and `message` you'll get a `data` like so

```json
{
  "code": "LOGGED_IN",
  "data": {
    "accessToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoiYWNjZXNzIiwiaXNzIjoiZG5kIiwic3ViIjoiMDczZDhmN2ItMDIyNS00NGRjLWE3NTMtOGJmNzYyYmVkMzc0IiwiZXhwIjoxNzA2NDQwNjExfQ.YQ6shB0HGGw9tN5jo6cBzqjoB4LxGlNadC52exF_Hm7UYfsbf8uB-u1Sq7ukgkIkkHw-eR0VLwmjNCWmWoF6tA",
    "refreshToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJrZXkiOiIkMmEkMTAkQmYwZ2FlUmxkOXRzWVEzbnBXSkFBZVlhdFlFZG9yWVVtRFl5VHl6TVA0a0lSRHdna1B3Y2EiLCJ0eXBlIjoicmVmcmVzaCIsImlzcyI6ImRuZCIsInN1YiI6IjA3M2Q4ZjdiLTAyMjUtNDRkYy1hNzUzLThiZjc2MmJlZDM3NCJ9.pv76u4p-kfAwGu8VPwzAKv5lGclrI85T2Uuu0kCT24hlfLRnjpU7iktgtlPujWuB_NVHxBKlvz_qkmyWeqLxlw",
    "user": {
      "id": "073d8f7b-0225-44dc-a753-8bf762bed374",
      "name": "Gonzalo",
      "profilePicture": "",
      "surname": "Aguado"
    }
  },
  "message": "Logged in correctly"
}
```

The system uses a refresh token and and access token, this is so you can revoke access everywhere by revoking the token
The refresh token never expires unless revoked
The access token expires every 15 minutes
This request is also accompanied by some simple user data for easier reference later

## Refresh token

To refresh your token send a `POST` request to `/refresh` with the `Authorization` containing the refresh key
If everything went OK you should get a 200 code and the usual `code` and `message` and you'll get an aditional `data` containing an `accessToken`

## Revoke token

To revoke your refresh token send a `POST` request to `/revoke` with the `Authorization` containing the access token
If everything went OK you should get a 200 code and the usual `code` and `message`

## Recover password

To recover your password send a `POST` request to `/recover-password-request` with the following JSON body

```json
{
  "email": "gonagutor@gmail.com"
}
```

If everything went OK you should get a 200 code and the usual `code` and `message` and an email should be sent containing a token used in the next step

## Redeem password recovery

To redeem a password recovery send a `POST` request to `/recover-password` with the following body (token being the token received via email in the previous request)

```json
{
  "password": "Testtest1@",
  "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoicmVjb3ZlciIsImlzcyI6ImRuZCIsInN1YiI6IjA3M2Q4ZjdiLTAyMjUtNDRkYy1hNzUzLThiZjc2MmJlZDM3NCIsImV4cCI6MTcwNjUyODk2MX0.E6zO0A4WmM-zexm8TdBuZQ3w-ps00o_bML80l6SZ19-gOAYfUpHkMZJdesN7wjjGsu1tnwtLsgKD6KY8H2BZBA"
}
```

If everything went OK you should get a 200 code and the usual `code` and `message` and your new password should be the password you sent in the body
