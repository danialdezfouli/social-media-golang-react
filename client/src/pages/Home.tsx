import { Spinner } from "components/elements/Spinner";
import Text from "components/elements/Text";
import PostItem from "components/timeline/post/Post";
import useHomeTimelineQuery from "connection/queries/useHomeTimelineQuery";

export default function Home() {
  const { data, isLoading } = useHomeTimelineQuery();

  return (
    <section className="home-timeline pb-32">
      {isLoading && (
        <div className="p-6 text-blue-600">
          <Spinner />
        </div>
      )}

      {data?.posts && data.posts.length === 0 && (
        <div className="py-10 px-2">
          <Text align="center" className="mt-2">
            برای مشاهده مطالب جدید افراد بیشتری را دنبال کنید
          </Text>
        </div>
      )}
      {data?.posts.map((post) => (
        <PostItem
          key={post.post_id}
          post={post}
          parent={post.parent_id ? data.parents[post.parent_id] : undefined}
        />
      ))}
    </section>
  );
}
