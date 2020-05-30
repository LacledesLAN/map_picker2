db.createUser(
    {
        user : "YourUserName",
        pwd : "YourPasswordHere", 
        roles: [
            {
                role : "readWrite", 
                db : "map-picker"
            }
        ]
    }
)