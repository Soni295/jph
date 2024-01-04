import { faker } from "@faker-js/faker";
import { User } from "../types";

export function createUser(userId: number): User {
  const firstName = faker.person.firstName()
  const username = faker.internet.userName({ firstName })
  const email = faker.internet.email({ firstName })

  const suite = faker.helpers.arrayElement(['Suite', 'Apt.'])
  const addressNumber = faker.number.int({ min: 0, max: 10_000 })

  const address = {
    street: faker.location.street(),
    suite: `${suite} ${addressNumber}`,
    city: faker.location.city(),
    zipcode: faker.number.int({ min: 1000, max: 7000 }),
    geo: {
      lat: faker.location.latitude(),
      lon: faker.location.longitude()
    }
  }

  return {
    id: userId,
    name: firstName,
    username,
    email,
    address,
  }
}

export function createUsers(countUsers: number): User[] {
  const users = []
  for (let i = 0; countUsers > i; i++) {
    const newId = users.length + 1
    const user = createUser(newId)
    users.push(user)
  }
  return users
}