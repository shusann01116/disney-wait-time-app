/** @type {import('next').NextConfig} */
const nextConfig = {
  async redirects() {
    return [
      {
        source: '/standby',
        destination: '/standby/tdl',
        permanent: true
      }
    ]
  }
}

module.exports = nextConfig
