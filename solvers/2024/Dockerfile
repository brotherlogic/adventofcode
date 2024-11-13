FROM ruby:2.3.1

RUN mkdir /solver
WORKDIR /solver
COPY Gemfile /solver
COPY Gemfile.lock /solver
RUN bundle config --global frozen 1
RUN bundle install --without development test
COPY lib /snip/lib
EXPOSE 8080
ENTRYPOINT [ "bundle", "exec"]
CMD ["server.rb"]