import useSWR from "swr";

const ENDPOINT_SUBSCRIBER = "http://localhost:9090";

const fetcher = (...args) => fetch(...args).then((res) => res.json());

const LatestTemperature = () => {
  const { data } = useSWR(`${ENDPOINT_SUBSCRIBER}/latest`, fetcher, {
    refreshInterval: 3000,
  });
  const latestTemp = data?.latesttemp || 0;
  const latestTime = data?.latesttime || "";
  return (
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
          color: parseInt(latestTemp, 10) > 30 ? "red" : "green",
          fontSize: 60,
          textAlign: "center",
        }}
      >
        {`${parseInt(latestTemp, 10)} C`}
      </div>
      <div tyle={{ textAlign: "center", alignItems: "center" }}>
        {latestTime}
      </div>
    </div>
  );
};

const MaxTemperature = () => {
  const { data } = useSWR(`${ENDPOINT_SUBSCRIBER}/max`, fetcher, {
    refreshInterval: 3000,
  });
  const maxTemp = data?.maxtemp || 0;
  return (
    <div
      style={{
        padding: 10,
        borde: "solid 2px #ddd",
        borderRadius: "5px",
        background: "#fff",
      }}
    >
      <div
        style={{ color: "red", fontSize: 20, textAlign: "center" }}
      >{`Max (${parseInt(maxTemp, 10)}C)`}</div>
    </div>
  );
};

const MinTemperature = () => {
  const { data } = useSWR(`${ENDPOINT_SUBSCRIBER}/min`, fetcher, {
    refreshInterval: 3000,
  });
  const minTemp = data?.mintemp || 0;
  return (
    <div
      style={{
        padding: 10,
        border: "solid 2px #ddd",
        borderRadius: "5px",
        background: "#fff",
      }}
    >
      <div
        style={{ color: "green", fontSize: 20, textAlign: "center" }}
      >{`Min (${parseInt(minTemp, 10)}C)`}</div>
    </div>
  );
};

const History = () => {
  const { data } = useSWR(`${ENDPOINT_SUBSCRIBER}`, fetcher, {
    refreshInterval: 3000,
  });
  const temperatures = data?.temperatures;
  return (
    <div>
      <h2>History</h2>
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
            <div style={{ textAlign: "center" }}>{temperature.timestamp}</div>
          </div>
        ))}
    </div>
  );
};

function App() {
  return (
    <div style={{ marginTop: 40 }}>
      <div style={{ width: 300, margin: "0 auto" }}>
        <h2>Latest</h2>
        <LatestTemperature />
        <MaxTemperature />
        <MinTemperature />
        <History />
      </div>
    </div>
  );
}

export default App;
