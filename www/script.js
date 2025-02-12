const d = window.document;

async function request(uri) {
  const response = await fetch(uri);
  if (!response.ok) {
    throw new Error(`request failed: ${response.status}`);
  }
  return await response.json();
}

function hydrate(resp) {
  const img = d.createElement('img');
  const container = getById('gifs');
  const title = getById('title');

  container.innerHTML = '';
  img.src = resp.data.images.original.url;
  title.innerText = resp.data.title;
  container.appendChild(img);
}

function getById(id) {
  return d.getElementById(id);
}
