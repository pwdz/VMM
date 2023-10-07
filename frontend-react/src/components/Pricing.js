import React, { useState, useEffect } from 'react';
import './Pricing.css';
import { fetchPricingData, updatePricingData } from '../services/apiService';

function Pricing() {
  const [cpuCost, setCpuCost] = useState({ id: 0, cost_per_unit: 0 });
  const [ramCost, setRamCost] = useState({ id: 0, cost_per_unit: 0 });
  const [hardDriveCost, setHardDriveCost] = useState({ id: 0, cost_per_unit: 0 });
  const [priceConfigs, setPriceConfigs] = useState([]);

  // Fetch pricing data when the component is loaded
  useEffect(() => {
    async function fetchData() {
      try {
        const pricingData = await fetchPricingData();
        // Update state with fetched pricing data
        setPriceConfigs(pricingData);
        // Find and set the CPU, RAM, and Hard Drive costs based on fetched data
        const cpuConfig = pricingData.find((config) => config.type === 'cpu');
        const ramConfig = pricingData.find((config) => config.type === 'ram');
        const hardDriveConfig = pricingData.find((config) => config.type === 'hdd');

        if (cpuConfig) {
          setCpuCost(cpuConfig);
        }

        if (ramConfig) {
          setRamCost(ramConfig);
        }

        if (hardDriveConfig) {
          setHardDriveCost(hardDriveConfig);
        }
      } catch (error) {
        console.error('Error fetching pricing data:', error);
      }
    }

    fetchData();
  }, []);

  const handleSave = async () => {
    try {
      // Prepare the updated priceConfigs array with ID
      const updatedPriceConfigs = [
        { id: cpuCost.id, type: 'cpu', cost_per_unit: cpuCost.cost_per_unit },
        { id: ramCost.id, type: 'ram', cost_per_unit: ramCost.cost_per_unit },
        { id: hardDriveCost.id, type: 'hdd', cost_per_unit: hardDriveCost.cost_per_unit },
      ];

      // Call the updatePricingData function to update pricing data on the server
      await updatePricingData(updatedPriceConfigs);
      console.log('Pricing data updated successfully');
    } catch (error) {
      console.error('Error updating pricing data:', error);
    }
  };

  return (
    <div className="pricing-container">
      <h1>Cost per CPU (IRT)</h1>
      <input
        type="number"
        value={cpuCost.cost_per_unit}
        onChange={(e) => setCpuCost({ ...cpuCost, cost_per_unit: parseInt(e.target.value, 10) })}
        className="input-field"
      />

      <h1>Cost per RAM (IRT)</h1>
      <input
        type="number"
        value={ramCost.cost_per_unit}
        onChange={(e) => setRamCost({ ...ramCost, cost_per_unit: parseInt(e.target.value, 10) })}
        className="input-field"
      />

      <h1>Cost per Hard Drive (GB)</h1>
      <input
        type="number"
        value={hardDriveCost.cost_per_unit}
        onChange={(e) => setHardDriveCost({ ...hardDriveCost, cost_per_unit: parseInt(e.target.value, 10) })}
        className="input-field"
      />

      <button className="save-button" onClick={handleSave}>
        Save
      </button>
    </div>
  );
}

export default Pricing;
