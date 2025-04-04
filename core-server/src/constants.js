// Database configuration constants
export const DB_CONNECTION_LIMIT = 10;
export const DB_NAME = process.env.DB_NAME || "my_project_db";
export const DB_USER = process.env.DB_USER || "postgres";
export const DB_PASSWORD = process.env.DB_PASSWORD || "password";
export const DB_HOST = process.env.DB_HOST || "localhost";
export const DB_PORT = process.env.DB_PORT || 5432;
export const DB_ACQUIRE_TIMEOUT = 10000;
export const DB_IDLE_TIMEOUT = 1800; 
export const QUERY_TIMEOUT = 60000; 

// Data constants
export const defaultProfilePic = "https://icon-library.com/images/anonymous-avatar-icon/anonymous-avatar-icon-25.jpg";

// Export everything
export default {
  DB_CONNECTION_LIMIT,
  DB_NAME,
  DB_USER,
  DB_PASSWORD,
  DB_HOST,
  DB_PORT,
  DB_ACQUIRE_TIMEOUT,
  DB_IDLE_TIMEOUT,
  QUERY_TIMEOUT,
  DEFAULT_PROFILE_PIC,
};
