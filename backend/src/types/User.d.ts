// This is the User type. It is used to represent a user in the database.
export type User = {
  id: number;
  username: string;
  email: string;
  password: string;
  created_at: Date;
  updated_at: Date;
};
