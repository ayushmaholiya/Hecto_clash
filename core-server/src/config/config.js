/*
    Configuration File

    This file contains the configuration settings for the application. 
    It retrieves values from environment variables or falls back to default values.

    Configuration Properties:
        port: number
            - The port on which the server will run.
            - Default: 3030 (if `process.env.PORT` is not set).

        dbHost: string
            - The hostname or IP address of the database server.
            - Default: "localhost" (if `process.env.DB_HOST` is not set).

        [add others]
 */
import{
    DB_CONNECTION_LIMIT,
    DB_NAME,
    DB_USER,
    DB_PASSWORD,
    DB_HOST,
    DB_PORT,
    DB_ACQUIRE_TIMEOUT,
    DB_IDLE_TIMEOUT,
    QUERY_TIMEOUT,
      } from "../constants.js"


        const conf = {
            // Server settings
            port: parseInt(process.env.PORT) || 3030,
        
            // Database settings
            dbHost: process.env.DB_HOST || "localhost",
            dbPort: parseInt(process.env.DB_PORT) || 5432,
            dbName: process.env.DB_NAME || "hecto_clash_db",
            dbUser: process.env.DB_USER || "postgres",
            dbPassword: process.env.DB_PASSWORD || "password",
            dbConnectionLimit: parseInt(process.env.DB_CONNECTION_LIMIT) || 10,
            dbIdleTimeout: parseInt(process.env.DB_IDLE_TIMEOUT) || 1800,
            dbAcquireTimeout: parseInt(process.env.DB_ACQUIRE_TIMEOUT) || 10000,
            queryTimeout: parseInt(process.env.QUERY_TIMEOUT) || 60000,
            
            // Application settings
            defaultProfilePic: process.env.DEFAULT_PROFILE_PIC || "https://example.com/default-avatar.jpg"
        };
        
        export default conf;