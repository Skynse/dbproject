var endpoint = "https://localhost:8080/api";

function getItem(name: string, id: string) {
  if (!name || !id) {
    throw new Error("Invalid arguments");
  }

  // if name, search by name
  // if id, search by id

  if (name) {
    return fetch(`${endpoint}/item?name=${name}`);
  }

  return fetch(`${endpoint}/item?id=${id}`);
}
