var dbUser = process.env.MONGO_USER;
var dbPassword = process.env.MONGO_PASSWORD;
db = db.getSiblingDB("db")


db.createUser({
    user: 'user',
    pwd: 'password',
    roles: [
      {
        role: 'dbOwner',
      db: 'db',
    },
  ],
}); 
db.createCollection("accounts");
db.createCollection("users");
db.createCollection("offices");
db.createCollection("roles");
db.createCollection("layers");