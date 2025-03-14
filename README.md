# Delivery Optimiser

## 1. Setup & Installation

There are two ways to run this project:
**1. Using Makefile** (Requires Go **1.24+**)  
**2. Using Docker** (Requires Docker installed)


### Run Locally Using Makefile
Ensure Go **1.24+** is installed, then run:
```
make build
make run
```

### Run Using Docker
Requires Docker installed.
Run the following command:
```./docker_run.sh```


## 2. Input & Output

### Input Files:
**input.json**: simulates an external service request for optimal delivery routing. The reviewer can modify this file to test different scenarios.
**input_validation.json**: defines validation rules for the input. This is maintained by the Delivery Optimiser server and should not be changed by the reviewer.

### Output Files (Generated in output/ directory):
#### Success Case:
**routes.json**: the computed optimal route response.
**delivery_instructions.txt**: step-by-step driver instructions.

#### Error Case:
**error.log**: captures execution errors if validation or processing fails.

## 3. Supported Algorithms
**Greedy**: At each step, the driver picks the next task (pickup/dropoff) based on the shortest travel time, considering restaurant preparation time when picking up. This provides a fast but not necessarily optimal solution.

**Brute Force**: Evaluates all possible order sequences, computing the total travel time for each, and selects the optimal sequence with the least total time. This guarantees the best route but is computationally expensive.

## 4. Workflow & Architecture

### Design Principles
**Encapsulation**: core.Run() encapsulates execution, keeping internal modules independent.
**Extendability**: optimiser allows plugging in new algorithms without modifying other modules.
**Modularity**: Internal packages handle distinct responsibilities for clean separation of concerns.

## Execution Flow:
The program starts execution via core.Run(), which orchestrates the process (using facade design pattern).
Parser loads input.json and input_validation.json.
Validation checks if input follows defined constraints.
Timegraph constructs travel time calculations between locations.
Optimiser selects the algorithm (Greedy, Brute Force) and computes the best route.
Output formats and writes results to the output/ directory.

## Internal Packages
**parser**: Loads and structures input data.
**validation**: Ensures input meets constraints.
**timegraph**: Builds time-based distance relationships using haversine algorithm.
**optimiser**: Encapsulates multiple algorithms and allows easy extension.
**output**: Generates and writes structured output.
**utils**: Utility functions for JSON, file handling etc.

## 5. Limitations & Future Enhancements
#### Logging & Debugging
Error logging can be structured better with proper bubbling for better traceability. A more detailed logging system including debug logs could improve visibility, but due to time constraints this was not implemented.

#### Testing
Due to time constraints test cases were not implemented. Various types of testing, including unit tests for individual packages and system wide integration tests could improve reliability and ensure correctness.

#### Dynamic Programming for Optimisation
A DP based approach can be added for better scalability. For this situation where each driver handles a limited number of restaurants (≤5), brute force is manageable.

#### Validation Ordering
The usual approach is to validate and structure the input at the beginning. Since Haversine distance is computed later in the time graph, the max distance validation is also done there. Ideally this validation should be moved earlier in the process.




	
	