package main

import (
	"testing"
)

func TestNeuronCreation(t *testing.T) {
	neuron := &Neuron{}

	if neuron == nil {
		t.Fatal("Expected neuron to be created")
	}

	if len(neuron.In) != 0 {
		t.Errorf("Expected 0 inputs, got %d", len(neuron.In))
	}

	if len(neuron.Out) != 0 {
		t.Errorf("Expected 0 outputs, got %d", len(neuron.Out))
	}
}

func TestNeuronIter(t *testing.T) {
	neuron := &Neuron{}
	result := neuron.Iter()

	if len(result) != 1 {
		t.Errorf("Expected 1 neuron from Iter(), got %d", len(result))
	}

	if result[0] != neuron {
		t.Error("Expected Iter() to return the neuron itself")
	}
}

func TestNeuronConnectTo(t *testing.T) {
	neuron1 := &Neuron{}
	neuron2 := &Neuron{}

	neuron1.ConnectTo(neuron2)

	if len(neuron1.Out) != 1 {
		t.Errorf("Expected 1 output connection, got %d", len(neuron1.Out))
	}

	if len(neuron2.In) != 1 {
		t.Errorf("Expected 1 input connection, got %d", len(neuron2.In))
	}

	if neuron1.Out[0] != neuron2 {
		t.Error("Expected neuron1.Out to contain neuron2")
	}

	if neuron2.In[0] != neuron1 {
		t.Error("Expected neuron2.In to contain neuron1")
	}
}

func TestNeuronLayerCreation(t *testing.T) {
	layer := NewNeuronLayer(5)

	if layer == nil {
		t.Fatal("Expected layer to be created")
	}

	if len(layer.Neurons) != 5 {
		t.Errorf("Expected 5 neurons, got %d", len(layer.Neurons))
	}
}

func TestNeuronLayerIter(t *testing.T) {
	layer := NewNeuronLayer(3)
	result := layer.Iter()

	if len(result) != 3 {
		t.Errorf("Expected 3 neurons from Iter(), got %d", len(result))
	}

	// Verify all returned neurons are pointers to layer neurons
	for i, neuron := range result {
		if neuron != &layer.Neurons[i] {
			t.Errorf("Expected neuron %d to match layer neuron", i)
		}
	}
}

func TestConnectTwoNeurons(t *testing.T) {
	neuron1 := &Neuron{}
	neuron2 := &Neuron{}

	Connect(neuron1, neuron2)

	if len(neuron1.Out) != 1 {
		t.Error("Expected neuron1 to have 1 output")
	}

	if len(neuron2.In) != 1 {
		t.Error("Expected neuron2 to have 1 input")
	}
}

func TestConnectNeuronToLayer(t *testing.T) {
	neuron := &Neuron{}
	layer := NewNeuronLayer(3)

	Connect(neuron, layer)

	if len(neuron.Out) != 3 {
		t.Errorf("Expected neuron to have 3 outputs, got %d", len(neuron.Out))
	}

	for i := 0; i < 3; i++ {
		if len(layer.Neurons[i].In) != 1 {
			t.Errorf("Expected layer neuron %d to have 1 input, got %d", i, len(layer.Neurons[i].In))
		}
	}
}

func TestConnectLayerToNeuron(t *testing.T) {
	layer := NewNeuronLayer(2)
	neuron := &Neuron{}

	Connect(layer, neuron)

	for i := 0; i < 2; i++ {
		if len(layer.Neurons[i].Out) != 1 {
			t.Errorf("Expected layer neuron %d to have 1 output, got %d", i, len(layer.Neurons[i].Out))
		}
	}

	if len(neuron.In) != 2 {
		t.Errorf("Expected neuron to have 2 inputs, got %d", len(neuron.In))
	}
}

func TestConnectLayerToLayer(t *testing.T) {
	layer1 := NewNeuronLayer(2)
	layer2 := NewNeuronLayer(3)

	Connect(layer1, layer2)

	// Each neuron in layer1 should have 3 outputs (one to each neuron in layer2)
	for i := 0; i < 2; i++ {
		if len(layer1.Neurons[i].Out) != 3 {
			t.Errorf("Expected layer1 neuron %d to have 3 outputs, got %d", i, len(layer1.Neurons[i].Out))
		}
	}

	// Each neuron in layer2 should have 2 inputs (one from each neuron in layer1)
	for i := 0; i < 3; i++ {
		if len(layer2.Neurons[i].In) != 2 {
			t.Errorf("Expected layer2 neuron %d to have 2 inputs, got %d", i, len(layer2.Neurons[i].In))
		}
	}
}

func TestMultipleConnections(t *testing.T) {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)

	// Verify complex connectivity
	if len(neuron1.Out) != 1+3 { // 1 to neuron2 + 3 to layer1
		t.Errorf("Expected neuron1 to have 4 outputs, got %d", len(neuron1.Out))
	}

	if len(neuron1.In) != 4 { // 4 from layer2
		t.Errorf("Expected neuron1 to have 4 inputs, got %d", len(neuron1.In))
	}
}

func TestCompositePatternConsistency(t *testing.T) {
	// Both Neuron and NeuronLayer implement NeuronInterface
	neuron := &Neuron{}
	layer := NewNeuronLayer(2)

	var ni1 NeuronInterface = neuron
	var ni2 NeuronInterface = layer

	if ni1 == nil || ni2 == nil {
		t.Error("Expected both to implement NeuronInterface")
	}
}
