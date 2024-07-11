var dbUser = process.env.MONGO_USER;
var dbPassword = process.env.MONGO_PASSWORD;
db = db.getSiblingDB("appDb")


db.createUser({
    user: dbUser,
    pwd: dbPassword,
    roles: [
      {
        role: 'dbOwner',
      db: 'appDb',
    },
  ],
}); 
db.createCollection("accounts");
db.createCollection("users");
db.createCollection("offices");
db.createCollection("roles");
db.createCollection("layers");