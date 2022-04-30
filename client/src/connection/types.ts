type PostTypesShape = "post" | "quote" | "repost" | "reply";

export interface IPost {
  post_id: number;
  user_id: number;
  profile_name: string;
  profile_image: string;
  profile_username: string;
  parent_id: number;
  post_type: PostTypesShape;
  content: string;
  favorites_count: number;
  replies_count: number;
  repost_count: number;
  quote_count: number;

  reposted: boolean;
  liked: boolean;

  created_at: string;
}

export interface IProfile {
  id: number;
  name: string;
  bio?: string;
  username: string;
  image: string;
  followed: boolean;
  official: boolean;
  following_count: number;
  followers_count: number;
}
