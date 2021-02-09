#include <iostream>

using namespace std;

#define MAX 10

class Queue{
    int queue[MAX], front, rear;

public:
    Queue(): front(-1), rear(-1){}
    void insert(int);
    int del();
    void display();
};

void Queue::insert(int value){
    if(rear == MAX-1){
        cout<<"The Queue is full !\n";
    }else{
        rear++;
        queue[rear] = value;
    }
}

int Queue::del(){
    if(front == rear){
        cout<<"The Queue is empty !\n";
        return -1;
    }else{
        front++;

        int value = queue[front];
        return value;
    }
}

void Queue::display(){
    if(front == rear){
        cout<<"The Queue is empty !\n";
    }else{
        for(int i=front+1;i<=rear;i++){
            cout<<queue[i]<<" ";
        }
        cout<<endl;
    }
}

int main(){
    int option, value;
    Queue q;

    do{
        cout<<"\n************** MENU **************\n"
        <<"1. Insert\n"
        <<"2. Delete\n"
        <<"3. Display\n"
        <<"4. Exit\n";

        cout<<"Enter you option: ";
        cin>>option;

        switch (option)
        {
        case 1:
            cout<<"Enter the value to be inserted: ";
            cin>>value;
            q.insert(value);
            break;
        
        case 2:
            value = q.del();
            if(value != -1){
                cout<<"The deleted element is: "<<value<<endl;
            }
            break;
        
        case 3:
            q.display();
            break;
        
        case 4:
            break;

        default:
            cout<<"Wrong option !\n";
            break;
        }
    }while(option != 4);

    if(option == 4){
        cout<<"\nThank you for using the program !\n"
        <<"Have a good day !\n";
    }
}

