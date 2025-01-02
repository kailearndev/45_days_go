<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Project Setup</title>
</head>
<body>
    <h1>Project Setup Instructions</h1>

    <h2>1. Install GORM</h2>
    <p>To install GORM, run the following command:</p>
    <pre><code>go get -u gorm.io/gorm</code></pre>
    <p>For more information, visit the official documentation: <a href="https://gorm.io/" target="_blank">https://gorm.io/</a></p>

    <h2>2. Connect to the Database</h2>
    <p>Use GORM to connect to your database. Example:</p>
    <pre><code>
dsn := "root:password@tcp(127.0.0.1:3306)/your_db?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
    panic("failed to connect database")
}
    </code></pre>

    <h2>3. Install Gin</h2>
    <p>To install Gin, run the following command:</p>
    <pre><code>go get -u github.com/gin-gonic/gin</code></pre>
    <p>For more information, visit the official GitHub repository: <a href="https://github.com/gin-gonic/gin" target="_blank">https://github.com/gin-gonic/gin</a></p>

    <h2>4. Install dotenv</h2>
    <p>To install the dotenv package, run the following command:</p>
    <pre><code>go get github.com/joho/godotenv</code></pre>
    <p>For more information, visit the official GitHub repository: <a href="https://github.com/joho/godotenv" target="_blank">https://github.com/joho/godotenv</a></p>

</body>
</html>
