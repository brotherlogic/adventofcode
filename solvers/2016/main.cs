using System;

namespace Solver2016 {
    class Program {
        const string Host = "0.0.0.0";

        public static void Main(string[] args) {
            // Build a server
            var server = new Server
            {
                Services = { GreetingService.BindService(new GreeterServiceImpl()) },
                Ports = { new ServerPort(Host, Port, ServerCredentials.Insecure) }
            };
        }
    }
}