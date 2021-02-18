// Step 1. Launch the Dockerfile
// Step 2. Create an IDE connection to the new database
// Step 3. Run the following queries in the console

use jetbrains
db.createCollection("ides")
db.ides.insertMany([
    {name: "GoLand", focus: "Go"},
    {name: "WebStorm", focus: "Web - Javascript/TypeScript"},
    {name: "PyCharm", focus: "Python"},
    {name: "IntelliJ IDEA Ultimate", focus: "Java & Everything"},
    {name: "Rider", focus: ".Net"},
])

SELECT name, focus FROM ides;

// Step 4. Use the (new) extractors from the list
// Step 5. Right click on a column value to get more editing features