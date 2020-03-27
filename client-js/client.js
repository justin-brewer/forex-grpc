var PROTO_PATH = '../converter/converter.proto';

var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });

var converter_proto = grpc.loadPackageDefinition(packageDefinition).converter;

function main() {
  var action = ""
  var client = new converter_proto.Converter('localhost:50051',
                                       grpc.credentials.createInsecure());

  if (process.argv.length >= 5) {
	source = process.argv[2];
	target = process.argv[3];
	amount = process.argv[4];
	action = "conversion";
  } else if (process.argv.length == 3)  {
	action = "list";
  } else {
	console.log('Usage: args = <source> <target> <amount>');
  }

  if (action === "conversion") {
    var request = {source: source, target: target, amount: parseFloat(amount)};
    client.getConversion(request, function(err, response) {
      console.log(`${amount} of ${source} = ${response.amount} of ${target}`);
    });
  } else if (action === "list") {
    var request = {request: ""};
    client.getCurrencyList(request, function(err, response) {
	    console.log("available currencies");
	    console.log(response.reply);
    });
  } else {
	  console.log("unsupported arguments");

  }
}

main();
