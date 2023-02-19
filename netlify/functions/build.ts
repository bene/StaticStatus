import { schedule } from "@netlify/functions";

const handler = schedule("* * * * *", async () => {
  const webhookUrl = process.env.BUILD_WEBHOOK;

  if (!webhookUrl) {
    throw new Error("No build webhook configured");
  }

  const res = await fetch(webhookUrl, {
    method: "POST",
  });

  return {
    statusCode: res.status,
  };
});

export { handler };
