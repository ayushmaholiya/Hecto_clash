import {Router} from "express"
import UserController from "../controllers/user.js";

const router = Router();

// User routes
router.post('/signup', UserController.Signup);
router.post('/login', UserController.Login);

export default router;