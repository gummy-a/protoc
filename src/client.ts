import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import path from "path";

const packageDefinition = protoLoader.loadSync(
  path.join(__dirname, "../message.proto"),
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
  }
);

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;
const messageProto = protoDescriptor.message;

async function main() {
  // Wait a bit for server to start
  await new Promise((resolve) => setTimeout(resolve, 1000));

  // Create client
  const client = new messageProto.StringService(
    "localhost:8080",
    grpc.credentials.createInsecure()
  )

  // Create message
  const message = {
    content: "Hello from client!",
    id: 42,
  };

  console.log(`Client sent - ID: ${message.id}, Content: ${message.content}`);

  // Call Echo method
  client.echo(message, (error: any, response: any) => {
    if (error) {
      console.error("Failed to call Echo:", error);
      process.exit(1);
    }

    if (response) {
      console.log(
        `Client received - ID: ${response.id}, Content: ${response.content}`
      );
    }

    process.exit(0);
  });
}

main();
