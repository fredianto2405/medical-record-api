CREATE FUNCTION emr_patient.generate_medical_record_number()
RETURNS VARCHAR AS $$
DECLARE
today DATE := current_date;
  new_number INT;
  formatted_number TEXT;
BEGIN
  LOCK TABLE emr_patient.medical_record_sequence IN EXCLUSIVE MODE;

  IF EXISTS (
    SELECT 1 FROM emr_patient.medical_record_sequence
    WHERE sequence_date = today
  ) THEN
UPDATE emr_patient.medical_record_sequence
SET last_number = last_number + 1
WHERE sequence_date = today
    RETURNING last_number INTO new_number;
ELSE
    INSERT INTO emr_patient.medical_record_sequence(sequence_date, last_number)
    VALUES (today, 1)
    RETURNING last_number INTO new_number;
END IF;

  formatted_number := LPAD(new_number::TEXT, 5, '0');
RETURN 'EMR-' || TO_CHAR(today, 'YYYYMMDD') || '-' || formatted_number;
END;
$$ LANGUAGE plpgsql;
