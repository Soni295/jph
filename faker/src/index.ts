import { createUsers } from "./create/user";
import { createPosts } from "./create/post";
import { saveData } from "./utils";

const users = createUsers(10)
const usersIds = users.map(user => user.id)
const posts = createPosts(100, usersIds)

const usersStr = JSON.stringify(users)
const postsStr = JSON.stringify(posts)
saveData('users', usersStr)
saveData('posts', postsStr)