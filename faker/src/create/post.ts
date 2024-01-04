import { faker } from "@faker-js/faker"
import { Post } from "../types"
import { getOneId } from "../utils"

export function createPost(userId: number, postId: number): Post {
  return {
    id: postId,
    userId: userId,
    title: faker.lorem.words(),
    body: faker.lorem.sentence({min:3, max:100})
  }
}

export function createPosts(countPosts: number, usersIds: number[]): Post[] {
  const posts = []
  for (let i = 0; countPosts > i; i++) {
    const userId = getOneId(usersIds)
    const newId = posts.length + 1
    const newPost = createPost(userId, newId)
    posts.push(newPost)
  }
  return posts
}