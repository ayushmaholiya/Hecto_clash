import {
  createUser,
  findUserByEmail
} from "../repositories/user.js";
import { generateToken } from "../utils/jwt.js";
import { defaultProfilePic } from "../constants.js";

const signupUser = async (email, username, age, profilePic = defaultProfilePic) => {
  try {
    const existingUser = await findUserByEmail(email);
    if (existingUser) {
      return {
        user: existingUser,
        message: "User already exists",
      };
    }

    const newUser = await createUser(email, username, age, profilePic);

    const token = generateToken({
      id: newUser.id,
      email: newUser.email,
      type: "user",
    });

    return {
      user: {
        id: newUser.id,
        email: newUser.email,
        username: newUser.username,
        age: newUser.age,
        profile_pic: newUser.profile_pic,
        current_rating: newUser.current_rating,
      },
      token,
    };
  } catch (error) {
    console.error("Signup Error:", error.message);
    throw new Error(error.message || "User signup failed");
  }
};

const loginUser = async (email) => {
  try {
    const user = await findUserByEmail(email);
    if (!user) {
      throw new Error("User not found. Please sign up.");
    }

    const token = generateToken({
      id: user.id,
      email: user.email,
      type: "user",
    });

    return {
      user: {
        id: user.id,
        email: user.email,
        username: user.username,
        age: user.age,
        profile_pic: user.profile_pic,
        current_rating: user.current_rating,
      },
      token,
    };
  } catch (error) {
    console.error("Login Error:", error.message);
    throw new Error(error.message || "User login failed");
  }
};

export { signupUser, loginUser };
