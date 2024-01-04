import fs from "node:fs"
import { faker } from "@faker-js/faker";

export const getOneId = (ids: number[]) => faker.helpers.arrayElement(ids)

export function saveData(entity: string, json: string) {
  const writeStream = fs.createWriteStream(`../src/compose/data/${entity}.json`)
  writeStream.write(json)
}