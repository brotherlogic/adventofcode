FROM gcr.io/google-appengine/aspnetcore:1.1

RUN mkdir /app
ADD . /app
WORKDIR /app

EXPOSE 8080
ENTRYPOINT ["dotnet", "2016.dll"]