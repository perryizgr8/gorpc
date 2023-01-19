import grpc
import procedures_pb2_grpc
import procedures_pb2

channel = grpc.insecure_channel("localhost:5000")
stub = procedures_pb2_grpc.PingStub(channel)
pong = stub.Ping(procedures_pb2.EchoRequest(msg="Ping!"))
print(pong)
