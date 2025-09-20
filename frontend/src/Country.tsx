function Country(props: any) {
  return (
    <div>
        <h1>{props.name} ({props.abbreviation})</h1>
        <h3>Capital: {props.capital}</h3>
        <h3>Continent: {props.continent}</h3>
    </div>
  );
}

export default Country;
