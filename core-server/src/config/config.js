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

const conf = {
	// server settings
	port: process.env.PORT || 3030,

	// database settings
	dbHost: process.env.DB_HOST || "localhost",
};

export default conf;
