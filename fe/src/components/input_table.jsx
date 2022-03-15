export default function InputTable() {
  return (
    <div className="container mt-5">
      <div className="row justify-content-md-center">
        <h1 className="col col-lg-4">Input Students</h1>
      </div>
      <div className="container">
        <div className="row justify-content-md-center">
          <form
            className="col-6 g-3"
            action="http://localhost:8080/students"
            method="POST"
          >
            <div className="mb-3">
              <label className="form-label">Full Name</label>
              <input
                name="full_name"
                type="text"
                className="form-control"
                id="full_name"
              />
            </div>
            <div className="mb-3">
              <label className="form-label">Extra ID</label>
              <input
                name="extra_id"
                type="number"
                className="form-control"
                id="extra_id"
              />
            </div>
            <button type="submit" className="btn btn-primary">
              Submit
            </button>
          </form>
        </div>
      </div>
    </div>
  );
}
