db = db.getSiblingDB("appDb");

db.createCollection("accounts");
print("Collection 'accounts' created");

db.createCollection("users");
print("Collection 'users' created");

db.createCollection("offices");
print("Collection 'offices' created");

db.createCollection("roles");
print("Collection 'roles' created");

db.createCollection("screens");
print("Collection 'layers' created");

print("init script finish");


