

export default function App() {
  return (
    <main className="min-h-screen bg-slate-950 text-white selection:bg-purple-500 selection:text-white overflow-x-hidden">
      {/* Background Gradients */}
      <div className="fixed inset-0 z-0 overflow-hidden pointer-events-none">
        <div className="absolute top-0 left-1/4 w-96 h-96 bg-purple-600/20 rounded-full blur-[128px]" />
        <div className="absolute bottom-0 right-1/4 w-96 h-96 bg-blue-600/20 rounded-full blur-[128px]" />
      </div>

      {/* Navbar */}
      <nav className="relative z-10 border-b border-white/10 bg-slate-950/50 backdrop-blur-xl">
        <div className="container mx-auto px-6 h-16 flex items-center justify-between">
          <div className="text-xl font-bold bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text text-transparent">
            Monorepo<span className="text-white">POC</span>
          </div>
          <div className="flex items-center gap-6 text-sm font-medium text-slate-400">
            <a href="#features" className="hover:text-white transition-colors">Features</a>
            <a href="#stack" className="hover:text-white transition-colors">Tech Stack</a>
            <a href="http://localhost:8081/health" target="_blank" className="hover:text-white transition-colors">API Health</a>
            <a
              href="http://localhost:8080"
              target="_blank"
              className="px-4 py-2 bg-white text-slate-950 rounded-full hover:bg-slate-200 transition-colors"
            >
              Login (Keycloak)
            </a>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="relative z-10 pt-32 pb-20 px-6 text-center">
        <div className="container mx-auto max-w-4xl">
          <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-white/5 border border-white/10 text-xs font-medium text-purple-300 mb-8">
            <span className="relative flex h-2 w-2">
              <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-purple-400 opacity-75"></span>
              <span className="relative inline-flex rounded-full h-2 w-2 bg-purple-500"></span>
            </span>
            v1.0.0 Now Live (React/Vite)
          </div>

          <h1 className="text-5xl md:text-7xl font-bold tracking-tight mb-8 bg-gradient-to-b from-white to-white/60 bg-clip-text text-transparent">
            Microservices Architecture <br />
            <span className="bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text text-transparent">Done Right.</span>
          </h1>

          <p className="text-lg md:text-xl text-slate-400 mb-12 max-w-2xl mx-auto leading-relaxed">
            A production-ready polyglot monorepo featuring Go services, React frontend,
            full observability stack, and OIDC authentication.
          </p>

          <div className="flex flex-col sm:flex-row items-center justify-center gap-4">
            <a
              href="#features"
              className="px-8 py-4 bg-purple-600 hover:bg-purple-500 text-white rounded-full font-medium transition-all hover:scale-105 shadow-lg shadow-purple-500/25"
            >
              Explore Features
            </a>
            <a
              href="http://localhost:16686"
              target="_blank"
              className="px-8 py-4 bg-white/5 hover:bg-white/10 border border-white/10 text-white rounded-full font-medium transition-all hover:scale-105 backdrop-blur-sm"
            >
              View Traces
            </a>
          </div>
        </div>
      </section>

      {/* Features Grid */}
      <section id="features" className="relative z-10 py-24 bg-slate-950/50">
        <div className="container mx-auto px-6">
          <h2 className="text-3xl font-bold text-center mb-16">Core Capabilities</h2>
          <div className="grid md:grid-cols-3 gap-8">
            <FeatureCard
              title="Polyglot Services"
              description="Go for high-performance backend services, Vite/React for modern frontend applications."
              icon="🚀"
            />
            <FeatureCard
              title="Full Observability"
              description="End-to-end tracing with Jaeger, metrics with Prometheus, and dashboards in Grafana."
              icon="📊"
            />
            <FeatureCard
              title="Secure Auth"
              description="Centralized identity management with Keycloak using OIDC and JWT validation."
              icon="🔒"
            />
          </div>
        </div>
      </section>

      {/* Stack Grid */}
      <section id="stack" className="relative z-10 py-24 border-t border-white/10">
        <div className="container mx-auto px-6">
          <h2 className="text-3xl font-bold text-center mb-16">Technology Stack</h2>
          <div className="flex flex-wrap justify-center gap-4 md:gap-8 opacity-70">
            {['React', 'Vite', 'Golang', 'Docker', 'PostgreSQL', 'Redis', 'Keycloak', 'Prometheus', 'Jaeger'].map((tech) => (
              <div key={tech} className="px-6 py-3 rounded-xl bg-white/5 border border-white/10 font-mono text-sm">
                {tech}
              </div>
            ))}
          </div>
        </div>
      </section>

      <footer className="relative z-10 py-8 border-t border-white/10 text-center text-slate-500 text-sm">
        <p>© 2024 Monorepo POC. Built with ❤️.</p>
      </footer>
    </main>
  )
}

function FeatureCard({ title, description, icon }: { title: string, description: string, icon: string }) {
  return (
    <div className="p-8 rounded-3xl bg-white/5 border border-white/10 hover:border-purple-500/50 transition-colors group">
      <div className="text-4xl mb-6 group-hover:scale-110 transition-transform duration-300">{icon}</div>
      <h3 className="text-xl font-bold mb-3 text-white group-hover:text-purple-400 transition-colors">{title}</h3>
      <p className="text-slate-400 leading-relaxed">
        {description}
      </p>
    </div>
  );
}
