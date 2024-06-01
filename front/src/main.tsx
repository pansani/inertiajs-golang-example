import { createRoot } from 'react-dom/client';
import { createInertiaApp } from '@inertiajs/inertia-react';
import { InertiaProgress } from '@inertiajs/progress';

InertiaProgress.init({
  delay: 250,
  color: '#29d',
  includeCSS: true,
  showSpinner: true
});

console.log("Iniciando aplicação Inertia");

const el = document.getElementById('app');

if (el) {
  console.log("Elemento com ID 'app' encontrado");
  if (el.dataset.page) {
    const initialPage = JSON.parse(el.dataset.page);
    console.log("Página inicial carregada", initialPage);

    createInertiaApp({
      page: initialPage,
      resolve: async (name: string) => {
        console.log(`Resolvendo componente: ${name}`);
        const module = await import(`./Pages/${name}`);
        console.log(`Componente ${name} carregado`);
        return module.default;
      },
      setup({ el, App, props }) {
        console.log("Configurando Inertia App");
        createRoot(el).render(<App {...props} />);
        console.log("Aplicação renderizada");
      },
    });
  } else {
    console.error("Atributo 'data-page' não encontrado no elemento com ID 'app'");
  }
} else {
  console.error("Elemento com ID 'app' não encontrado");
}
