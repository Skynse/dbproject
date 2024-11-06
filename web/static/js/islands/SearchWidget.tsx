import { useSignal } from "@preact/signals";

export default function ItemSearch(nameSearchFilter, idSearchFilter) {
  // Signals for search, results, and loading state
  const searchName = useSignal(""); // Name search input
  const searchId = useSignal(""); // ID search input
  const results = useSignal([]);
  const loading = useSignal(false);

  // Function to handle the search based on the search criteria (name or ID)
  const searchItems = async () => {
    setLoading(true);
    let endpoint = "localhost:8080/api/"; // Adjust for your backend or API endpoint

    // Check if searching by name or ID
    const query = searchName.value.trim() || searchId.value.trim();

    // Only proceed if there is a query (either name or ID)
    if (!query) {
      alert("Please enter a search term.");
      setLoading(false);
      return;
    }

    const response = await getItem(searchName.value, searchId.value);
    const data = await response.json();

    // Ensure the results are correctly populated (the actual data is in the 'v' property)
    setResults(data); // Store data in results signal
    setLoading(false);
  };

  // Function to fetch items based on search criteria
  const getItem = async (name, id) => {
    let url = `http://localhost:8080/api/item?name=${name}`;
    if (id) {
      url = `http://localhost:8080/api/item?id=${id}`;
    }
    return fetch(url);
  };

  // Function to set the loading state
  const setLoading = (state) => {
    loading.value = state;
  };

  // Function to set the results
  const setResults = (data) => {
    console.log(data);
    results.value = [data]; // Store data in results signal
  };

  return (
    <div className="min-h-screen bg-gray-100 flex justify-center items-center p-4">
      <div className="max-w-3xl w-full bg-white rounded-lg shadow-lg p-6">
        <h1 className="text-3xl font-bold text-center text-blue-600 mb-6">
          Item Search
        </h1>

        {/* Search by Name */}
        <div className="mb-4">
          <label
            htmlFor="search-name"
            className="block text-gray-700 font-medium mb-2"
          >
            Search by Name:
          </label>
          <input
            id="search-name"
            type="text"
            value={searchName.value}
            onInput={(e) => (searchName.value = e.target.value)}
            className="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Enter item name"
          />
        </div>

        {/* Search by ID */}
        <div className="mb-6">
          <label
            htmlFor="search-id"
            className="block text-gray-700 font-medium mb-2"
          >
            Search by ID:
          </label>
          <input
            id="search-id"
            type="text"
            value={searchId.value}
            onInput={(e) => (searchId.value = e.target.value)}
            className="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Enter item ID"
          />
        </div>

        {/* Search Button */}
        <div className="flex justify-center mb-6">
          <button
            onClick={searchItems}
            className="w-full md:w-1/2 py-3 bg-blue-600 text-white font-semibold rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            Search
          </button>
        </div>

        {/* Loading State */}
        {loading.value ? (
          <div className="text-center text-blue-500">Loading...</div>
        ) : (
          <ul className="space-y-4">
            {results.value != null ? (
              results.value.map((item) => (
                <li
                  key={item.iId} // Use the correct ID property
                  className="p-4 bg-gray-50 border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow duration-200"
                >
                  <strong className="text-lg text-blue-600">
                    {item.Iname}
                  </strong>
                  <p className="text-sm text-gray-500">ID: {item.iId}</p>
                  <p className="text-sm text-gray-500">Price: ${item.Sprice}</p>
                  <p className="text-sm text-gray-500">
                    Description: {item.Idescription}
                  </p>
                </li>
              ))
            ) : (
              <p className="text-center text-gray-500">No results found</p>
            )}
          </ul>
        )}
      </div>
    </div>
  );
}
