export interface User {
  id: number,
  name: string,
  username: string,
  email: string,
  address: Address
}

interface Address {
  street: string
  suite: string
  city: string
  zipcode: number
  geo: Geolocalitation
}

interface Geolocalitation {
  lat: number
  lon: number
}