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

  var client = new converter_proto.Converter('localhost:50051',
                                       grpc.credentials.createInsecure());

  if (process.argv.length >= 5) {
	source = process.argv[2];
	target = process.argv[3];
	amount = process.argv[4];
  } else {
	console.log('Usage: args = <source> <target> <amount>')
  }
  var request = {source: source, target: target, amount: parseFloat(amount)};

  client.getConversion(request, function(err, response) {
    console.log(`${amount} of ${source} = ${response.amount} of ${target}`);
  });

}

main();
