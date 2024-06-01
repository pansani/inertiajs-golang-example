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
  const initialPage = JSON.parse(el.dataset.page || '');

  createInertiaApp({
    page: initialPage,
    resolve: name => import(`./Pages/${name}`).then(module => module.default),
    setup({ el, App, props }) {
      createRoot(el).render(<App {...props} />);
    },
  });
} else {
  console.error("Elemento com ID 'app' não encontrado");
}

