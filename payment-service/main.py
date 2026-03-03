from fastapi import FastAPI, HTTPException, status
from pydantic import BaseModel
from typing import List, Optional

app = FastAPI()

payments = []
id_counter = 1

class Payment(BaseModel):
    orderId: int
    amount: float
    method: str

@app.get("/payments")
def get_payments():
    """Returns all payments """
    return payments

@app.post("/payments/process", status_code=status.HTTP_201_CREATED)
def process_payment(payment: Payment):
    """Process a payment """
    global id_counter
    payment_data = payment.dict()
    payment_data["id"] = id_counter
    payment_data["status"] = "SUCCESS"
    
    payments.append(payment_data)
    id_counter += 1
    return payment_data

@app.get("/payments/{payment_id}")
def get_payment(payment_id: int):
    """Get payment status by ID """
    payment = next((p for p in payments if p["id"] == payment_id), None)
    
    if payment is None:
        raise HTTPException(status_code=404, detail="Payment not found")

    return payment