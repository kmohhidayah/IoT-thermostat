import useSWR from "swr";

const ENDPOINT_SUBSCRIBER = "http://localhost:4000";

const getData = async () => {
  const response = await fetch(ENDPOINT_SUBSCRIBER);
  return await response.json();
};

function App() {
  const { data } = useSWR(ENDPOINT_SUBSCRIBER, getData);
  const latestTemp = data?.latest;
  const temperatures = data?.temperatures;
  console.log(temperatures);

  return (
    <div style={{ marginTop: 40 }}>
      <div style={{ width: 300, margin: "0 auto" }}>
        {latestTemp && (
          <div
            style={{
              padding: 10,
              border: "solid 2px #ddd",
              borderRadius: "5px",
              background: "#fff",
            }}
          >
            <div
              style={{
                color: parseInt(latestTemp.temp, 10) > 30 ? "red" : "green",
                fontSize: 60,
                textAlign: "center",
              }}
            >
              {`${parseInt(latestTemp.temp, 10)} C`}
            </div>
            <div tyle={{ textAlign: "center" }}>{latestTemp.timestamp}</div>
          </div>
        )}
        <div>
          {temperatures &&
            temperatures.map((temperature) => (
              <div
                key={`${temperature.id + temperature.temp}`}
                style={{
                  padding: 10,
                  background: "#fff",
                  border: "solid 2px #ddd",
                }}
              >
                <div
                  style={{
                    fontSize: 25,
                    textAlign: "center",
                    color: parseInt(temperature.temp) > 30 ? "red" : "green",
                  }}
                >
                  {`${parseInt(temperature.temp, 10)} C`}
                </div>
                <div style={{ textAlign: "center" }}>
                  {temperature.timestamp}
                </div>
              </div>
            ))}
        </div>
      </div>
    </div>
  );
}

export default App;
