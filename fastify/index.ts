import Fastify from 'fastify'

const fastify = Fastify({logger: false})

fastify.get('/', function (request, reply) {
  reply.send("Hello world (node)!")
})

fastify.listen({ port: 3000 }, function (err, address) {
  if (err) {
    fastify.log.error(err)
    process.exit(1)
  }
})