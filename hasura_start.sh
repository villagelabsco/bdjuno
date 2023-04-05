docker run -d -p 8080:8080 \
  -e HASURA_GRAPHQL_DATABASE_URL=postgres://wurts:@localhost:5432/wurts \
  -e HASURA_GRAPHQL_ENABLE_CONSOLE=true \
  -e HASURA_GRAPHQL_UNAUTHORIZED_ROLE="anonymous" \
  hasura/graphql-engine:latest