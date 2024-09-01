###----- Deps -----###
FROM node:22-bookworm-slim AS deps

WORKDIR /app

COPY package.json ./package.json
COPY package-lock.json ./package-lock.json

RUN npm ci

###----- Builder -----###
FROM node:22-bookworm-slim AS builder

WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules
COPY . .

RUN npm run build

###----- Production -----###
FROM node:22-bookworm-slim AS production

WORKDIR /app

ARG UID=1001
ARG GID=1001
ARG USERNAME="next"
ARG GROUPNAME="next"

RUN mkdir .next && \
    groupadd -g ${GID} ${GROUPNAME} && \
    useradd --create-home -u ${UID} -g ${GID} ${USERNAME} && \
    chown -R ${UID}:${GID} .next

COPY --from=builder --chown=${USERNAME}:${GROUPNAME} /app/public ./public
COPY --from=builder --chown=${USERNAME}:${GROUPNAME} /app/.next/standalone ./
COPY --from=builder --chown=${USERNAME}:${GROUPNAME} /app/.next/static /static

USER next

ARG NODE_ENV=production
ENV NODE_ENV=${NODE_ENV}

EXPOSE ${PORT}

CMD [ "node", "server.js" ]
