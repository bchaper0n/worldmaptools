function CountryCard(props: any) {
  return (
    <div className="col">
      <div className="card" style={{width: "18rem"}}>
          <img src={"data:image/svg+xml;base64," + props.flag} className="card-img-top" alt={props.name + " Flag"}></img>
          <div className="card-body">
            <h5 className="card-title">{props.name} ({props.abbreviation})</h5>
            <p className="card-text">{props.continent}</p>
            <p className="card-text">Capital: {props.capital}</p>
            <a href="#" className="btn btn-primary">more info</a>
          </div>
        </div>
    </div>
  );
}

export default CountryCard;
