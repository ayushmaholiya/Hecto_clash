import { asyncHandler, apiResponse, apiError } from "../utils/index.js";
import { signupUser, loginUser } from "../services/auth.js";
import { defaultProfilePic } from "../constants.js";

/*
 * @desc    User Signup
 * @route   POST /api/v1/user/signup
 * @access  Public
 */
const Signup = asyncHandler(async (req, res) => {
  const { email, username, age, profilePic, auth0Token } = req.body;

  if (!email || !username || !age || !auth0Token) {
    throw new apiError(400, "Email, Username, Age, and Auth0 Token are required");
  }

  const { user, token } = await signupUser(email, username, age, profilePic || defaultProfilePic, auth0Token);

  return res
    .status(201)
    .json(new apiResponse(201, "User signed up successfully.", { user, token }));
});

/*
 * @desc    User Login
 * @route   POST /api/v1/user/login
 * @access  Public
 */
const Login = asyncHandler(async (req, res) => {
  const { email, auth0Token } = req.body;

  if (!email || !auth0Token) {
    throw new apiError(400, "Email and Auth0 Token are required");
  }

  const { user, token } = await loginUser(email, auth0Token);

  return res
    .status(200)
    .json(new apiResponse(200, "User logged in successfully.", { user, token }));
});

export default { Signup, Login };
