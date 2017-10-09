# ---------------------------------------------
# Yarn
# ---------------------------------------------
FROM kkarczmarczyk/node-yarn as yarn

WORKDIR /app
COPY ./package.json ./
COPY ./yarn.lock ./

RUN yarn

# ---------------------------------------------
# Bower
# ---------------------------------------------
FROM node as bower

WORKDIR /app
COPY ./bower.json ./
COPY ./.bowerrc ./

RUN npm install --global bower

RUN bower install --allow-root

# ---------------------------------------------
# Final Build FROM NODE
# ---------------------------------------------
FROM node

WORKDIR /app

COPY --from=yarn /app/node_modules ./node_modules
COPY --from=bower /app/public/libs /app/public/libs

COPY ./ ./

CMD ["bash"]