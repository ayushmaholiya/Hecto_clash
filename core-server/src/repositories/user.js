import getDBConnection from "../utils/dbPool.js";

// Create User 
const createUser = async (username, email, age, profilePic = null, currentRating = 400) => {
  let conn;
  try {
    conn = await getDBConnection();
    if (!conn) {
      throw new Error(`Couldn't connect to the database.`);
    }

    const result = await conn.query(
      `INSERT INTO users (username, email, age, profile_pic, current_rating) 
       VALUES ($1, $2, $3, $4, $5) RETURNING id`, 
      [username, email, age, profilePic, currentRating]
    );

    return {
      id: result.rows[0].id,
      username,
      email,
      age,
      profile_pic: profilePic,
      current_rating: currentRating,
    };
  } catch (err) {
    console.error(err);
    throw err;
  } finally {
    if (conn) {
      conn.release(); 
    }
  }
};

// Find User by Email
const findUserByEmail = async (email) => {
  let conn;
  try {
    conn = await getDBConnection();
    if (!conn) {
      throw new Error(`Couldn't connect to the database.`);
    }

    const result = await conn.query(
      "SELECT * FROM users WHERE email = $1", 
      [email]
    );

    return result.rows[0] || null;
  } catch (err) {
    console.error(err);
    return null;
  } finally {
    if (conn) {
      conn.release(); 
    }
  }
};

// Find User by ID
const findUserById = async (id) => {
  let conn;
  try {
    conn = await getDBConnection();
    if (!conn) {
      throw new Error(`Couldn't connect to the database.`);
    }

    const result = await conn.query(
      "SELECT * FROM users WHERE id = $1",
      [id]
    );

    return result.rows[0] || null;
  } catch (err) {
    console.error(err);
    return null;
  } finally {
    if (conn) {
      conn.release();
    }
  }
};

export { createUser, findUserByEmail, findUserById };