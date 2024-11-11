import { useSignal } from "@preact/signals";

export default function Dashboard() {
  const searchName = useSignal("");
  const searchId = useSignal("");
  const results = useSignal([]);
  const loading = useSignal(false);
  const itemName = useSignal("");
  const itemId = useSignal("");
  const itemPrice = useSignal(0.0);
  const itemDescription = useSignal("");
  const apiUrl = "http://localhost:8020/api";
  const selectedItem = useSignal({});

  // CRUD Operations

  // Fetch All Items
  const fetchItems = async () => {
    setLoading(true);
    const response = await fetch(`${apiUrl}/items`);
    const data = await response.json();
    results.value = Array.isArray(data) ? data : [data];
    setLoading(false);
  };

  // Add Item
  const addItem = async () => {
    const newItem = {
      Iname: itemName.value,
      Sprice: Number(itemPrice.value),
      Idescription: itemDescription.value,
    };
    console.log("Sending item data:", JSON.stringify(newItem));
    await fetch(`${apiUrl}/item/insert`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(newItem),
    });
    fetchItems(); // Refresh the list
  };

  const search = async () => {
    setLoading(true);
    if (searchName.value && searchId.value) {
      const response = await fetch(
        `${apiUrl}/item/search?name=${searchName.value}&id=${searchId.value}`,
      );
      const data = await response.json();
      results.value = Array.isArray(data) ? data : [data];
    } else if (searchName.value) {
      const response = await fetch(
        `${apiUrl}/item/search?name=${searchName.value}`,
      );
      const data = await response.json();
      results.value = Array.isArray(data) ? data : [data];
    } else if (searchId.value) {
      const response = await fetch(
        `${apiUrl}/item/search?id=${searchId.value}`,
      );
      const data = await response.json();
      results.value = Array.isArray(data) ? data : [data];
    }
    setLoading(false);
  };

  // Update Item
  const updateItem = async () => {
    console.log(selectedItem.value);
    const updatedItem = {
      iId: selectedItem.value.iId,
      Iname: selectedItem.value.Iname,
      Sprice: Number(selectedItem.value.Sprice),
      Idescription: selectedItem.value.Idescription,
    };
    await fetch(`${apiUrl}/item`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(updatedItem),
    });
    fetchItems(); // Refresh the list

    showPopup.value = false;
  };

  // Delete Item
  const deleteItem = async (id) => {
    await fetch(`${apiUrl}/item`, {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ iId: id }),
    });
    fetchItems(); // Refresh the list
  };

  // Helper function to manage loading state
  const setLoading = (state) => {
    loading.value = state;
  };

  const showPopup = useSignal(false);

  const togglePopup = () => {
    showPopup.value = !showPopup.value;
  };

  return (
    <div className="min-h-screen bg-gray-100 flex justify-center items-center p-4">
      <div className="max-w-4xl w-full bg-white rounded-lg shadow-lg p-6">
        <h1 className="text-3xl font-bold text-center text-blue-600 mb-6">
          Item Dashboard
        </h1>

        {/* Search Form */}
        {/* Search Form */}
        <div className="mb-4">
          <h2 className="text-xl font-semibold text-blue-600">Search Form</h2>
          <div className="grid grid-cols-2 gap-4 mt-2">
            <input
              type="text"
              placeholder="Search by ID"
              value={searchId.value}
              onInput={(e) => (searchId.value = e.target.value)}
              className="p-3 border rounded"
            />
            <input
              type="text"
              placeholder="Search by Name"
              value={searchName.value}
              onInput={(e) => (searchName.value = e.target.value)}
              className="p-3 border rounded"
            />
          </div>
          <button
            onClick={search}
            className="mt-4 bg-blue-600 text-white py-2 px-4 rounded"
          >
            Search
          </button>
        </div>

        {/* Form to Add Item */}
        <div className="mb-4">
          <h2 className="text-xl font-semibold text-blue-600">Add Item</h2>
          <div className="grid grid-cols-2 gap-4 mt-2">
            <input
              type="text"
              placeholder="Item Name"
              value={itemName.value}
              onInput={(e) => (itemName.value = e.target.value)}
              className="p-3 border rounded"
            />
            <input
              type="text"
              placeholder="Price"
              value={itemPrice.value}
              onInput={(e) => (itemPrice.value = e.target.value.toString())}
              className="p-3 border rounded"
            />
            <input
              type="text"
              placeholder="Description"
              value={itemDescription.value}
              onInput={(e) => (itemDescription.value = e.target.value)}
              className="p-3 border rounded"
            />
          </div>
          <button
            onClick={addItem}
            className="mt-4 bg-green-600 text-white py-2 px-4 rounded"
          >
            Add Item
          </button>
        </div>

        {/* Fetch and Display All Items */}
        <div className="mb-4">
          <button
            onClick={fetchItems}
            className="bg-blue-600 text-white py-2 px-4 rounded"
          >
            Fetch All Items
          </button>
        </div>

        {/* Loading Indicator */}
        {loading.value ? (
          <div className="text-center text-blue-500">Loading...</div>
        ) : (
          <table className="min-w-full bg-white border border-gray-200 rounded-lg">
            <thead>
              <tr>
                <th className="px-4 py-2 border">ID</th>
                <th className="px-4 py-2 border">Name</th>
                <th className="px-4 py-2 border">Price</th>
                <th className="px-4 py-2 border">Description</th>
                <th className="px-4 py-2 border">Actions</th>
              </tr>
            </thead>
            <tbody>
              {results.value.map((item) => (
                <tr key={item.iId} className="text-center border">
                  <td className="px-4 py-2 border">{item.iId}</td>
                  <td className="px-4 py-2 border">{item.Iname}</td>
                  <td className="px-4 py-2 border">${item.Sprice}</td>
                  <td className="px-4 py-2 border">{item.Idescription}</td>
                  <td className="px-4 py-2 border space-x-2">
                    <button
                      onClick={() => deleteItem(item.iId)}
                      className="bg-red-600 text-white py-1 px-2 rounded"
                    >
                      Delete
                    </button>
                    <button
                      //  show edit form
                      onClick={() => {
                        selectedItem.value = item;
                        togglePopup();
                      }}
                      className="bg-yellow-500 text-white py-1 px-2 rounded"
                    >
                      Edit
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
      {showPopup == true ? (
        <div className="fixed top-0 left-0 w-full h-full bg-gray-900 bg-opacity-50 flex justify-center items-center">
          <div className="bg-white p-6 rounded-lg shadow-lg">
            <h2 className="text-xl font-semibold text-blue-600">Edit Item</h2>
            <div className="grid grid-cols-2 gap-4 mt-2">
              <input
                type="text"
                placeholder="Item Name"
                value={selectedItem.value.Iname}
                onInput={(e) => (selectedItem.value.Iname = e.target.value)}
                className="p-3 border rounded"
              />
              <input
                type="text"
                placeholder="Price"
                value={selectedItem.value.Sprice}
                onInput={(e) => (selectedItem.value.Sprice = e.target.value)}
                className="p-3 border rounded"
              />
              <input
                type="text"
                placeholder="Description"
                value={selectedItem.value.Idescription}
                onInput={(e) =>
                  (selectedItem.value.Idescription = e.target.value)
                }
                className="p-3 border rounded"
              />
            </div>
            <button
              onClick={updateItem}
              className="mt-4 bg-yellow-500 text-white py-2 px-4 rounded"
            >
              Update Item
            </button>
            <button
              onClick={togglePopup}
              className="mt-4 bg-red-600 text-white py-2 px-4 rounded"
            >
              Close
            </button>
          </div>
        </div>
      ) : null}
    </div>
  );
}
