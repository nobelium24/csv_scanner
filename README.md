# Student CSV Sorter

## Overview
This Go program reads a CSV file containing student records, sorts them by their matriculation numbers, and then saves the sorted data into separate CSV files categorized by the students' state of origin.

## Features
- Reads student data from a CSV file.
- Groups students by their state.
- Sorts students within each state by their `MATRIC_NUM`.
- Generates separate CSV files for each state.
- Handles errors like missing files and incorrect formats.

## Installation
1. Install [Go](https://go.dev/doc/install) if you haven't already.
2. Clone this repository or copy the `main.go` file to your local directory.
3. Navigate to the project directory:
   ```sh
   cd path/to/your/project
   ```

## Usage
1. Compile and run the program with a CSV file as an argument:
   ```sh
   go run main.go <csv_file>
   ```
   Example:
   ```sh
   go run main.go ENROLMENT_LIST_2023_2024.csv
   ```
2. The program processes the file and creates separate CSV files named after each state.

## Expected Output
- If `ENROLMENT_LIST_2023_2024.csv` contains students from different states, you will see output like:
  ```sh
  Created file: OGUN.csv
  Created file: LAGOS.csv
  Created file: KANO.csv
  Processing complete.
  ```
- These files will contain sorted student data for each state.

## Error Handling
- **File Not Found**: If the provided CSV file doesn't exist, you'll see:
  ```sh
  Error reading CSV file: open <filename>: no such file or directory
  ```
- **Invalid Data Format**: If the CSV structure is incorrect, the program will return an error and stop execution.
- **File Creation Error**: If a file cannot be created, an error message will be displayed, but the program will continue processing other states.

## Notes
- Ensure the CSV file has a proper header and follows the expected format.
- The program ignores case differences in state names, but be cautious of typos.
- If a state field contains invalid data (e.g., empty or "Select State"), an unexpected file name may be created.

## License
This project is open-source and can be modified or redistributed under the MIT License.

